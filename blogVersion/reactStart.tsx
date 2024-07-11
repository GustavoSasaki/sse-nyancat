import { useEffect, useState } from "react";

// <start>
// <show>
export function Nyancat() {

	const [flavourUrl, setFlavourUrl] = useState("vanilla")

	useEffect(() => {
		const sseSource = new EventSource("http://184.72.221.34:8080/events?stream=flavour");

		sseSource.onmessage = (e) => {
			setFlavourUrl(e.data);
		};

		window.onbeforeunload = function () {
			sseSource.close();
		};

		//ensures close when exit page
		sseSource.onerror = () => sseSource.close()
	}, [])

	return (
		<img src={flavourUrl} />
	)
}
// </show>

