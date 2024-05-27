import { CollectionConfig } from 'payload/types';

export const Users: CollectionConfig = {
	slug: 'users',
	auth: {
		useAPIKey: true,
	},
	admin: {
		useAsTitle: 'email',
	},
	fields: [],
};
