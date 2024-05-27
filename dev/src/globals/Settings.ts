import { GlobalConfig } from "payload/types";

export const Settings: GlobalConfig = {
	slug: 'settings',
	fields: [
		{
			name: 'siteName',
			type: 'text',
			label: 'Site Name',
		},
	]
}
