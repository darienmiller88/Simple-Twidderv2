<script lang="ts">
  import {type Tweed} from "./lib/types"
  import { Moon } from 'svelte-loading-spinners';
  import "./tweed.scss"

  let isLoading:    boolean = false
  let name:         string = ""
  let tweedContent: string = ""
  let tweeds:       Tweed[] = [
    {
      name: "Darien",
      content: "First tweed",
      date: new Date()
    },
    {
      name: "Dalton",
      content: "dreads",
      date: new Date()
    },
    {
      name: "Vicky",
      content: "spam!",
      date: new Date()
    }
  ]

  function sendTweed(){
    let tweed: Tweed = {
      name: name,
      content: tweedContent,
      date: new Date()
    }

    tweeds = [...tweeds, tweed]
    name = ""
    tweedContent = ""
    isLoading = !isLoading
  }
</script>

<main>
  <div class="title twidder-color">Twidder</div>

  {#if !isLoading}
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
  {/if}

  {#if isLoading}
    <div class="loading">
      <Moon size="200" color="#1DA1F2" unit="px" duration="0.75s" />
    </div>
  {/if}

  <div class="tweeds">
    {#each tweeds as tweed}
      <div class="tweed">
        <div class="name" id={tweed.date.toString() + tweed.name.at(0)}>{tweed.name}</div>
        <div class="content" id={tweed.date.toString() + tweed.name.at(0)}>{tweed.content}</div>
        <div class="date" id={tweed.date.toString() + tweed.name.at(0)}>{tweed.date.toString()}</div>
      </div>
    {/each}
  </div>
</main>

<style lang="scss">
  
</style>
