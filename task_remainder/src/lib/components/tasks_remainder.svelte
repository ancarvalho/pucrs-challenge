<script lang="ts">
  import { Daytime } from "../utils/day_time";
  import Trash from "../utils/icons/trash.svelte";

  export let dayTime: Daytime;

  import { delete_task, tasks, transform_into_tasks, filter_by_daytime, mark_as_finished } from "../store/tasks";

  transform_into_tasks()

  // let task;

  let filtered_tasks: typeof $tasks;
  tasks.subscribe((val) => {
    filtered_tasks = filter_by_daytime(dayTime, val);
  });

  
</script>

<div class="flex items-center flex-col text-black dark:text-white">
  <p class="text-center text-xl font-bold pb-8">
    {#if dayTime == Daytime.Day}
      DIA
    {:else}
      NOITE
    {/if}
  </p>

  <ul class="flex flex-col gap-2 text-lg font-medium items-center">
    {#each filtered_tasks as task}
      <li class={task.finished ? "flex gap-2 line-through" : "flex gap-2"}>
        {task.name} - {task.time}
        <button on:click={() => mark_as_finished(task.id)}>
          <Trash />
        </button>
      </li>
    {/each}
  </ul>
</div>
