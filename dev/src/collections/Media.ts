import { CollectionConfig } from "payload/types";
import { slateEditor } from '@payloadcms/richtext-slate'
import path from 'path';

export const Media: CollectionConfig = {
	slug: 'media',
	upload: {
		staticURL: '/media',
		staticDir: path.resolve(__dirname, '../../../media'),
	},
	fields: [
		{
			name: 'alt',
			type: 'text',
			required: true,
		},
		{
			name: 'caption',
			type: 'richText',
			editor: slateEditor({
				admin: {
					elements: ['link'],
				},
			}),
		},
	],
}
