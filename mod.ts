import { createBot, startBot } from "discord";
import { bgBlue } from "colors";

import { settings } from "@settings";
import { addDailyRecord, computeDailyScoreboardPoint } from "@usecases/dailyBot.ts";

console.log(bgBlue("Starting discord bot..."));

const token = settings.DISCORD_TOKEN;
const botId = BigInt(settings.BOT_ID);


const baseBot = createBot({
  token,
  botId,
  intents: ["Guilds", "GuildMessages"],
  events: {
    ready() {
      console.log("Successfully connected to gateway");
    },
    async messageCreate(bot, message) {
      const isDailyChannel = message.channelId === BigInt(settings.DAILY_CHANNEL);
      if (!isDailyChannel) return;

      const content = message.content.toLocaleLowerCase();
      const isDailyMessage = content.includes("o que eu fiz") && content.includes("o que vou fazer");
      if (!isDailyMessage) return;

      const userId = message.member!.id.toString();
      const name = message.member?.nick ?? message.tag;

      await addDailyRecord({ name, userId });

      const { error, data: dailyScoreboard } = await computeDailyScoreboardPoint(userId, name);
      if (error != null) {
        console.error(error);
        return;
      }

      console.log({ dailyScoreboard });
    },
  },
});

await startBot(baseBot);
