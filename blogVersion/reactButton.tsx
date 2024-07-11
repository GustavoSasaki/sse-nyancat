import { useEffect, useState } from "react";

// <start>
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
		<div>
			<img src={flavourUrl} />
			{/* // <show> */}
			<button
				onClick={async () => await fetch('http://184.72.221.34:8080/change-flavour')}>
				<p >Change flavour</p>
			</button>
			{/* // </show> */}
		</div>
	)
}

