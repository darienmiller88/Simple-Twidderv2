<script lang="ts">
  import {type Tweed} from "./lib/types"
  import DarkModeToggle from "./components/DarkModeToggle.svelte";
  import axios, { type AxiosResponse } from 'axios';

  import { Moon } from 'svelte-loading-spinners';
  import "./tweed.scss"
  import { onMount } from "svelte";

  let isDarkMode:   boolean = true
  let isLoading:    boolean = true
  let name:         string = ""
  let tweedContent: string = ""
  let URL         : string = window.location.hostname == "localhost" ? "http://localhost:8080/api/v1/tweeds" : "https://twidderapi.fly.dev/api/v1/tweeds"
  let tweeds:       Tweed[] = []

  onMount(() => getTweeds())

  async function sendTweed(){
    let tweed: Tweed = {
      id: "",
      name: name,
      content: tweedContent,
      created_at: new Date()
    }

    try {
      const tweedResponse: Tweed = await axios.post(URL, tweed)

      isLoading = true
      console.log(tweedResponse)
      tweeds = [tweed, ...tweeds]
      isLoading = false
    } catch (error) {
      console.log("err:", error)
    }

    name = ""
    tweedContent = ""
  }

  async function getTweeds(){
    try {
        const response: AxiosResponse<Tweed[]>  = await axios.get(URL)
        
        tweeds = response.data.reverse()
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
      <input placeholder="Enter name" bind:value={name} />
    </div>
    <br />
    <div class="text-container form-item">
      <label for="tweed" class="twidder-color">Tweed</label><br />
      <textarea placeholder="Enter tweed" rows="5" bind:value={tweedContent} />
    </div>
    <button>Send Tweed!</button>
  </form> 

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
    </div>
  {/if}

</main>

<style lang="scss">
  
</style>
