import { error } from "@sveltejs/kit";
import type { PageLoad } from "./$types";

export const load = (({params}) => {
	if (params.slug === "hello-world") {
		return {
			title: "Hello World!",
			content: "Welcome to our blog."
		};
	}

	throw error(404, "Not Found");
}) satisfies PageLoad;