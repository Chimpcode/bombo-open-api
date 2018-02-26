import React, { Component } from "react";
import InputText from "./components/InputText";
import CoolButton from "./components/CoolButton";
import styled, { css } from "styled-components";

const WrapContainer = styled.div`
	display: grid;
	grid-template-columns: 15% 70% 15%;
	grid-template-rows: 110px 80px 35px 71px 20px 100%;
`;

const Container = styled.div`
	grid-column: 2 / 3;
	grid-row: 2 / 3;
	display: flex;
`;

const ContainerField = styled.div`
	margin-right: 24px;
	padding: 0;
	${props =>
		props.flex &&
		css`
			flex: ${props.flex};
		`};
`;

const ContainerTitle = styled.span`
	grid-column: 2 / 3;
	grid-row: 0 / 1;

	font-family: "Quicksand", sans-serif;
	font-size: 28px;
	color: #5b5b5b;
	text-align: center;
	margin-top: 34px;
`;

const ContainerButton = styled.div`
	grid-column: 2 / 3;
	grid-row: 4 / 5;
`;

class App extends Component {
	render() {
		return (
			<WrapContainer>
				<ContainerTitle>WORKER PANEL</ContainerTitle>

				<Container>
					<ContainerField flex={3}>
						<InputText
							name="URL"
							onChange={e => console.log(e.target.value)}
						/>
					</ContainerField>

					<ContainerField flex={2}>
						<InputText
							name="Name"
							onChange={e => console.log(e.target.value)}
						/>
					</ContainerField>

					<ContainerField flex={1}>
						<InputText
							name="Type"
							onChange={e => console.log(e.target.value)}
						/>
					</ContainerField>

					<ContainerField flex={1}>
						<InputText
							name="Period"
							onChange={e => console.log(e.target.value)}
						/>
					</ContainerField>
				</Container>

				<ContainerButton>
					<CoolButton>CREATE NEW WORKER</CoolButton>
				</ContainerButton>
			</WrapContainer>
		);
	}
}

export default App;
