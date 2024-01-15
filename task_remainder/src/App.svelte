<script lang="ts">
  import ProfileSidebar from "./lib/components/profile_sidebar.svelte";
  import TasksRemainder from "./lib/components/tasks_remainder.svelte";
  import AddTask from "./lib/components/add_task.svelte";
  import Add from "./lib/utils/icons/add.svelte";

  let showModal = false;

  function handleModal() {
    showModal = !showModal;
  }


</script>

<main
  class="relative flex p-4 items-center justify-center bg-slate-100 dark:bg-slate-800 h-screen"
>
  <section class="min-w-[33%]">
    <ProfileSidebar />
  </section>
  <section class="min-w-[66%] flex flex-col items-center">
    <div class="flex items-center justify-center pb-8 gap-4">
      <p class="text-2xl font-bold">Lista de Tarefas</p>

      <button on:click={handleModal}>
        <Add />
      </button>
      <!-- <AddTask display={createDialogHidden} /> -->
    </div>

    <div class="flex justify-around w-full">
      <div class="min-w-[50%]">
        <TasksRemainder isNight={false} />
      </div>

      <div class="min-w-[50%]">
        <TasksRemainder isNight={true} />
      </div>
    </div>
  </section>

  {#if showModal}
    <!-- <svelte:window on:click={closeModal} /> -->
    <AddTask
      on:close={handleModal}
    />
  {/if}

</main>
