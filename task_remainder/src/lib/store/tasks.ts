import { derived, writable } from 'svelte/store'
import type { Task } from '../types/task'
import { Daytime } from '../utils/day_time'

const dummy_tasks = ["Escovar os dentes - 08:00", "Tirar o Lixo - 09:00", "Escovar os dentes - 08:00", "Tirar o Lixo - 09:00", "Escovar os dentes - 08:00", "Tirar o Lixo - 09:00",]



export function transform_into_tasks() {
  const dummytasks = dummy_tasks.map((t, i) => {
    console.log(i)
    let [task_name, time] = t.split("-")
    const task: Task = {
      id: i,
      name: task_name.trim(),
      time: time.trim(),
      finished: false,
      dayTime: i % 2 == 0 ? Daytime.Day : Daytime.Night
    }
    return task
  })
  tasks.set(dummytasks)
}

export function delete_task(id: number) {
  tasks.update((t) => t.filter((ta) => ta.id != id))
}

export function mark_as_finished(id: number) {
  console.log("Mark")
  tasks.update((t) => {
    const task_id = t.findIndex((ta) => ta.id === id)
    t[task_id] = {...t[task_id], finished: true}
    return t
  })
}

export function filter_by_daytime(dayTime: Daytime, tasks: Task[]) {
  return tasks.filter((t) => t.dayTime == dayTime)
}

export function create_task(task: Task) {
  tasks.update((t) => {
    t.push(task)
    return t
    
  })
}


export const tasks = writable<Task[]>()