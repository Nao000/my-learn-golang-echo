/**
 * 参考: https://nextjs.org/docs/pages/building-your-application/data-fetching/get-static-props
 */
export async function loadPosts() {

	const dummy = '{"id":1, "title":"THIS IS TITLE", "body":"THIS IS BODY TEXT.THIS IS BODY TEXT."}';

	return await JSON.parse(dummy);
}

export async function loadEmployees() {

	const res = await fetch('http://localhost/accessdb');
	const data = await res.json();

	return data;
}