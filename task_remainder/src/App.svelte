<script lang="ts">
  import ProfileSidebar from "./lib/components/profile_sidebar.svelte";
  import TasksRemainder from "./lib/components/tasks_remainder.svelte";
  import { Daytime } from "./lib/utils/day_time";
</script>

<main
  class="flex p-4 items-center justify-center bg-slate-100 dark:bg-slate-800 h-screen"
>
  <section class="min-w-[33%]">
    <ProfileSidebar />
  </section>
  <section class="min-w-[66%] flex flex-col items-center">
    <p class="text-2xl font-bold pb-8">Lista de Tarefas</p>

    <div class="flex justify-around w-full">
      <div class="min-w-[50%]">
        <TasksRemainder dayTime={Daytime.Day} />
      </div>

      
      <div class="min-w-[50%]">
        <TasksRemainder dayTime={Daytime.Night} />
      </div>
    </div>
  </section>
</main>
