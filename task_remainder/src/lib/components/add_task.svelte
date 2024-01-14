<script lang="ts">
    import { writable } from "svelte/store";

  export let display: boolean
  import { create_task } from "../store/tasks";
    import type { Task } from "../types/task";

  let new_task = writable<Task>()

 
 function handle_create_task(e: Event) {
  e.preventDefault()
 }


</script>


<div class={display ? "absolute" : "hidden"}>
  <form on:submit={handle_create_task}>
    <input type="text" bind:value={$new_task.name}>
    <input type="text" bind:value={$new_task.time}>
  </form>
</div>