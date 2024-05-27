import { mongooseAdapter } from '@payloadcms/db-mongodb';
import { webpackBundler } from '@payloadcms/bundler-webpack';
import { slateEditor } from '@payloadcms/richtext-slate';
import { buildConfig } from 'payload/config';

import {Users} from './collections/Users';
import {Posts} from './collections/Posts';
import {Settings} from "./globals/Settings";

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
	db: mongooseAdapter({
		url: process.env.DATABASE_URI,
	}),
});
