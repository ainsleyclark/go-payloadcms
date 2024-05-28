import {postgresAdapter} from '@payloadcms/db-postgres'
import { webpackBundler } from '@payloadcms/bundler-webpack';
import { slateEditor } from '@payloadcms/richtext-slate';
import { buildConfig } from 'payload/config';

import {Users} from './collections/Users';
import {Posts} from './collections/Posts';
import {Settings} from "./globals/Settings";
import {Media} from "./collections/Media";

export default buildConfig({
	admin: {
		user: Users.slug,
		bundler: webpackBundler(),
	},
	editor: slateEditor({}),
	collections: [
		Users,
		Posts,
		Media,
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
