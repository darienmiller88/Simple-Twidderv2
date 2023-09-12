<script lang="ts">
  import {type Tweed} from "./lib/types"
  import DarkModeToggle from "./components/DarkModeToggle.svelte";
  import axios, { type AxiosResponse } from 'axios';
  // import { IconHeart } from '@tabler/icons-svelte';

  import { Moon } from 'svelte-loading-spinners';
  import "./tweed.scss"
  import { onMount } from "svelte";

  let isDarkMode:    boolean = true
  let isLoading:     boolean = true
  let name:          string = ""
  let tweedContent:  string = ""
  let tweeds:        Tweed[] = []

  let isAtBottom:    boolean = false
  let finished:      boolean = false
  let skip:          number = 0
  let limit:         number = 5

  let isRateLimited:  boolean = false
  let rateLimitError: string = "Stop sending so much fucking tweeds!" 

  let URL:           string = window.location.hostname == "localhost" ? "http://localhost:8080/api/v1/tweeds" : "https://twidderapi.fly.dev/api/v1/tweeds"
  $:  if(isAtBottom && !finished){
        (async () => {
          const response = await axios.get(URL + `/tweed-range/${skip}/${limit}/desc`)

          if (response.data.meta.has_more) {
            skip += limit
            tweeds = [...tweeds, ...response.data.tweeds]
          }else{
            finished = true
          }
        })()
      }

  onMount(() => getTweeds())

  window.onscroll = () => {
    isAtBottom = window.innerHeight + window.scrollY >= document.body.offsetHeight
  }

  async function sendTweed(){
    let tweed: Tweed = {
      id: "",
      name: name,
      content: tweedContent,
      created_at: new Date()
    }

    try {
      const tweedResponse: Tweed = await axios.post(URL, tweed)

      console.log(tweedResponse)
      isLoading = true
      tweeds = [tweed, ...tweeds]
      isLoading = false
    } catch (error) {
      error.message
      console.log("err:", error)
    }

    name = ""
    tweedContent = ""
  }

  async function getTweeds(){
    try {
        const response = await axios.get(URL + `/tweed-range/${skip}/${limit}/desc`)
        
        tweeds = [...tweeds, ...response.data.tweeds]
        skip += limit
        isLoading = false
      } catch (error) {
        console.log("err;", error)
      }
  }

  function changeColorTheme(){
    isDarkMode = !isDarkMode

    if (!isDarkMode) {
        document.body.classList.add("light")
    } else {
        document.body.classList.remove("light")
    }
  }
</script>

<main>
  <div class="title twidder-color">Twidder</div>

  <DarkModeToggle isDarkMode={isDarkMode} changeColorTheme={changeColorTheme} />

  <form on:submit|preventDefault={sendTweed}>
    <div class="name-container form-item">
      <label for="name" class="twidder-color">Name</label><br />
      <input placeholder="Enter name" bind:value={name} minlength="3" maxlength="30" required/>
    </div>
    <br />
    <div class="text-container form-item">
      <label for="tweed" class="twidder-color">Tweed</label><br />
      <textarea placeholder="Enter tweed" rows="5" bind:value={tweedContent} minlength="2" maxlength="100" required/>
    </div>
    <button>Send Tweed!</button>
  </form> 

  {#if isRateLimited}
    <div class="rate-limit-error">
      {rateLimitError}
    </div>
  {/if}

  {#if isLoading}
    <div class="loading">
      <Moon size="200" color="#1DA1F2" unit="px" duration="0.75s" />
    </div>
  {:else}
    <div class="tweeds">
      {#each tweeds as tweed}
        <div class={`tweed ${isDarkMode ? "dark" : "light"}`} id={tweed.id}>
          <div class="name"   >{tweed.name}</div>
          <div class="content">{tweed.content}</div>
          <div class="date"   >{tweed.created_at.toString()}</div>
        </div>
      {/each}
      {#if isAtBottom && !finished}
        <div class="pagination-loading">Loading...</div>
      {/if}
    </div>
  {/if}

</main>

<style lang="scss">
  
</style>
