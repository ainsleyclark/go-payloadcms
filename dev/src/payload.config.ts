import {postgresAdapter} from '@payloadcms/db-postgres'
import { webpackBundler } from '@payloadcms/bundler-webpack';
import { slateEditor } from '@payloadcms/richtext-slate';
import { buildConfig } from 'payload/config';

import {Users} from './collections/Users';
import {Posts} from './collections/Posts';
import {Settings} from "./globals/Settings";

console.log('DATABASE_URI', process.env.DATABASE_URI);

export default buildConfig({
	admin: {
		user: Users.slug,
		bundler: webpackBundler(),
	},
	editor: slateEditor({}),
	collections: [
		Users,
		Posts,
	],
	globals: [
		Settings,
	],
	plugins: [],

	db: postgresAdapter({
		pool: {
			connectionString: process.env.DATABASE_URI,
		}
	}),
});
