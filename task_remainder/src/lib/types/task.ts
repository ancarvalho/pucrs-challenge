import type { Daytime } from "../utils/day_time"

export type Task = {
  id: number,
  name: string,
  time: string,
  dayTime: Daytime,
  finished:  boolean
}