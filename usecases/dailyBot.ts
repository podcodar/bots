import { client } from "@repository/supabase.ts";
import { DailyRecord } from "@entities";

export function addDailyRecord(dailyRecord: DailyRecord) {
  return client.from("daily_record").insert([dailyRecord]);
}

export function countLastWeekActivity() {
  // TODO: count dailies records from last week
  const today = new Date();
  const last7Days = new Date(Date.now() - 7 * 24 * 60 * 60 * 1000);

  return client
    .from("daily_record")
    .select("*", { count: "exact" })
    .lt("created_at", today)
    .gte("created_at", last7Days);
}

export function getScoreboardByUserId(userId: BigInt) {
  return client.from("daily_scoreboard").select("*").eq("userId", userId);
}

export function updateUserScoreboard(userId: BigInt, points: number, currentStreak: number) {
  return client.from("daily_scoreboard").update({ points, currentStreak }).eq("userId", userId);
}
