import { CollectionConfig } from "payload/types";

export const Posts: CollectionConfig = {
	slug: 'posts',
	fields: [
		{
			name: 'title',
			label: 'Title',
			type: 'text',
			required: true,
		},
		{
			name: 'content',
			label: 'Content',
			type: 'textarea',
		},
	],
}
