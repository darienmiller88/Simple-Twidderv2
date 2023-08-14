<script lang="ts">
  import {type Tweed} from "./lib/types"
  let name:  string = ""
  let tweedContent: string = ""
  let tweeds: Tweed[] = []

  function sendTweed(){
    let tweed = {
      name,
      tweedContent
    }

    console.log(tweed)
    name = ""
    tweedContent = ""
  }
</script>

<main>
  <div class="title twidder-color">Twidder</div>
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

  <div class="tweeds">

  </div>
</main>

<style lang="scss">
  .title{
    text-align: center;
    font-size: 70px;
    margin: 10px 0px;
    color: var(--twidder-blue);
  }

  form{
    text-align: center;
    width: 90%;
    margin: auto;

    .text-container{
      padding: 10px 0px;
            
      label{
        font-size: 20px;
        color: var(--twidder-blue);
      }

      textarea{
        width: 100%;
        font-size: 20px;
      }
    }

    .name-container{
      label{
        font-size: 20px;
        color: var(--twidder-blue);
      }
  
      input{
        width: 100%;
        font-size: 20px;
      }
    }

    button{
      width: 80%;
      padding: 15px 0px;
      font-size: 20px;

      border: 2px white solid;
      border-radius: 10px;

      color: white;
      background-color: var(--twidder-blue);
      transition: 0.5s;

      &:hover{
        cursor: pointer;
        background-color: var(--twidder-blue-transparent);
      }

      &:active{
        transform: translateY(-5px);
      }
    }

    @media only screen and (min-width: 768px) {
      width: 80%;
    }

    @media only screen and (min-width: 1200px) {
      width: 50%;
    }
  }

  .tweeds{
    border: 2px solid black;
    margin: auto;
    width: 90%;
    margin-top: 30px;

    @media only screen and (min-width: 768px) {
      width: 80%;
    }

    @media only screen and (min-width: 1200px) {
      width: 50%;
    }
  }

  .form-item{
    text-align: left;
    width: 80%;
    margin: auto;
    // text-align: left;
  }
</style>
