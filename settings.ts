import 'load-dotenv';

interface Settings {
	DISCORD_TOKEN: string;
	BOT_ID: string;
	DAILY_CHANNEL: string;
	SUPABASE_URL: string;
	SUPABASE_TOKEN: string;
}

// @ts-ignore
export const settings: Settings = {
	DISCORD_TOKEN: Deno.env.get('DISCORD_TOKEN') ?? '',
	BOT_ID: Deno.env.get('BOT_ID') ?? '',
	DAILY_CHANNEL: Deno.env.get('DAILY_CHANNEL') ?? '',
	SUPABASE_URL: Deno.env.get('SUPABASE_URL') ?? '',
	SUPABASE_TOKEN: Deno.env.get('SUPABASE_TOKEN') ?? '',
};
