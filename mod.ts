import { createBot, startBot } from "discord";
import { bgBlue } from "colors";

import { settings } from "@settings";
import { DailyRecord, DailyScoreboard } from "@entities";

console.log(bgBlue("Starting discord bot..."));

const token = settings.DISCORD_TOKEN;
const botId = BigInt(settings.BOT_ID);

const baseBot = createBot({
  token,
  botId,
  intents: ['Guilds', 'GuildMessages'],
  events: {
    ready() {
      console.log("Successfully connected to gateway");
    },
    messageCreate(bot, message) {
      const isDailyChannel = message.channelId === BigInt(settings.DAILY_CHANNEL);
      if (!isDailyChannel) return;

      const content = message.content.toLocaleLowerCase();
      const isDailyMessage =
        content.includes("o que fiz") &&
        content.includes("o que vou fazer");

      if (!isDailyMessage) return;

      const userId = message.member!.id;
      const name = message.member?.nick ?? message.tag;
      const dailyRecord: DailyRecord = {
        name,
        userId,
        date: new Date().toISOString(),
      }
      // TODO: load from database
      const dailyScoreboard: DailyScoreboard = {
        name,
        userId,
        points: 0,
        currentStreak: 0,
      }

      console.log(`content: ${message.content}`);
      console.log({message})
      // Process the message with your command handler here
    },
  },
});

await startBot(baseBot);
