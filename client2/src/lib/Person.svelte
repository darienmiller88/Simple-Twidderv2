<script lang="ts">
    import { type Person } from "./types"
    export let person: Person
</script>

<main>
    <div class="person"> 
        <div>{person.name}</div>
        <div>{person.age}</div>
    </div>
</main>

<style lang="scss">
    .person{
      border: 2px solid black;
      border-radius: 10px;
      width: fit-content;
      padding: 20px;
      background-color: rgba(100, 180, 200);
      transition: 0.4s;

      &:hover{
        cursor: pointer;
        background-color: rgba(100, 180, 200, 0.3);
      }
    }
</style>