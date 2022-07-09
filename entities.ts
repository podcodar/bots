export interface DailyRecord {
  date: string;
  userId: BigInt;
  name: string;
}

export interface DailyScoreboard {
  userId: BigInt;
  name: string;
  points: number;
  currentStreak: number;
}
