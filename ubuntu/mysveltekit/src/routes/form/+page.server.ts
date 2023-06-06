import type { Actions } from "./$types";

export const actions = {
	// (event) として受け取るとその他のデータも取得可能
	default: async ({ cookies, request }) => {
		const data = await request.formData();
		const myinput = data.get('myinput');
		console.dir({data, myinput});

	}
} satisfies Actions;