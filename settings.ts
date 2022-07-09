import { config } from "dotenv";

interface Settings {
  DISCORD_TOKEN: string;
  BOT_ID: string;
  DAILY_CHANNEL: string;
}

// @ts-ignore
export const settings: Settings = config();