export interface DailyRecord {
  userId: string;
  name: string;
}

export interface DailyScoreboard extends DailyRecord {
  points: number;
  currentStreak: number;
}
