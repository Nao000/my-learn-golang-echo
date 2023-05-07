type Props = {
	exampleArg: string
}

export default function Footer({ exampleArg }: Props) {
	return (

		<footer>{exampleArg} this is sample footer.</footer>
	);
}