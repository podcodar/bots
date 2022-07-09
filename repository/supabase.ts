import { createClient } from 'supabase';
import { settings } from '@settings';

const projectUrl = settings.SUPABASE_URL;
const token = settings.SUPABASE_TOKEN;

export const client = createClient(projectUrl, token);
