export interface DailyRecord {
  userId: string;
  name: string;
  id: string;
  created_at: string;
}

export type CreateDailyRecord = Omit<DailyRecord, 'id' | 'created_at'>;

export interface DailyScoreboard extends Omit<DailyRecord, 'created_at'> {
  points: number;
  currentStreak: number;
}

export type CreateDailyScoreboard = Omit<DailyScoreboard, 'id' | 'created_at'>;
