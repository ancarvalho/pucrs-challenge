<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import { writable } from "svelte/store";
  import type { Task } from "../types/task";

  import { create_task } from "../store/tasks";

  const task = writable<Task>({ isNight: false } as Task);

  const dispatch = createEventDispatcher();

  function handle_create_task(e: Event) {
    console.log(JSON.stringify($task));

    e.preventDefault();
    if ($task.name && $task.time) {
      create_task($task);
    }
    
    close();
  }

  function close() {
    dispatch("close");
  }

  let value: string;
</script>

<div class="absolute bg-slate-500">
  <div class="relative">
    <button
      class="absolute top-1 right-1 bg-slate-300 p-2 rounded-full"
      on:click={close}>X</button
    >
  </div>

  <div class="w-full max-w-xs">
    <form on:submit={handle_create_task} class="rounded px-8 pt-6 pb-8 mb-4">
      <div class="pb-4">
        <label
          class="block text-gray-700 dark:text-white text-sm font-bold mb-2"
          for="name"
        >
          Name
        </label>
        <input
          bind:value={$task.name}
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          id="name"
          type="text"
          placeholder="Name"
        />
      </div>
      <div class="pb-4">
        <label
          class="block text-gray-700 dark:text-white text-sm font-bold mb-2"
          for="time"
        >
          Time
        </label>
        <input
          bind:value={$task.time}
          class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          id="time"
          type="text"
          placeholder="Time"
        />
      </div>

      <label class="relative inline-flex items-center me-5 cursor-pointer">
        <input
          bind:checked={$task.isNight}
          type="checkbox"
          class="sr-only peer"
        />
        <div
          class="w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-focus:ring-4 peer-focus:ring-orange-300 dark:peer-focus:ring-orange-800 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-orange-500"
        ></div>
        <span class="ms-3 text-sm font-medium text-gray-900 dark:text-gray-300">
          Night Task ?
        </span>
      </label>

      <div class="flex justify-center mt-4">
        <button
          type="submit"
          class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
        >
          Salvar
        </button>
      </div>
    </form>
  </div>
</div>
