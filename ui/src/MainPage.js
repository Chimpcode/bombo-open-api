import React, { Component } from "react";
import InputText from "./components/InputText";
import CoolButton from "./components/CoolButton";
import styled, { css } from "styled-components";
import WorkerCard from "./components/WorkerCard";
import { observer } from "mobx-react";
import ContentLoader from "react-content-loader";
import axios from "axios";

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

const ContainerCards = styled.div`
	grid-column: 2 / 3;
	grid-row: 5 / 6;
`;

@observer
class MainPage extends Component {
	constructor(props) {
		super(props);
		this.state = { loading: true };
		this.newWorker = {};
	}

	fetchWorks = () => {
		axios
			.get("http://127.0.0.1:8500/manager/get-works", {
				auth: {
					username: "bregymr",
					password: "malpartida1"
				}
			})
			.then(response => {
				console.log(response.data.data);
				this.props.store.workers = response.data.data;

				this.setState({ loading: false });
			});
	};

	componentWillMount() {
		this.fetchWorks();
	}

	render() {
		const { workers } = this.props.store;
		const workersList = workers.map(work => {
			return (
				<WorkerCard
					data={work}
					key={work.url}
					// onRefresh={() => {
					// 	this.fetchWorks();
					// }}
				/>
			);
		});

		console.log(workersList);
		return (
			<WrapContainer>
				<ContainerTitle>WORKER PANEL</ContainerTitle>

				<Container>
					<ContainerField flex={3}>
						<InputText
							name="URL"
							onChange={e =>
								(this.newWorker.url = e.target.value)
							}
						/>
					</ContainerField>

					<ContainerField flex={2}>
						<InputText
							name="Name"
							onChange={e =>
								(this.newWorker.name = e.target.value)
							}
						/>
					</ContainerField>

					<ContainerField flex={1}>
						<InputText
							name="Type"
							onChange={e =>
								(this.newWorker.type = e.target.value)
							}
						/>
					</ContainerField>

					<ContainerField flex={1}>
						<InputText
							name="Period"
							onChange={e =>
								(this.newWorker.period = e.target.value)
							}
						/>
					</ContainerField>
				</Container>

				<ContainerButton>
					<CoolButton
						onClick={() =>
							this.props.store.createNewWork(this.newWorker)
						}
					>
						CREATE NEW WORK
					</CoolButton>
				</ContainerButton>

				<ContainerCards>
					{this.state.loading ? (
						<ContentLoader
							height={200}
							speed={2}
							primaryColor={"#f3f3f3"}
							secondaryColor={"#ecebeb"}
						>
							<rect
								x="0"
								y="23.05"
								rx="5"
								ry="5"
								width="400"
								height="54"
							/>

							<rect
								x="0.12"
								y="82.05"
								rx="4"
								ry="4"
								width="400"
								height="5.01"
							/>
						</ContentLoader>
					) : (
						workersList
					)}
				</ContainerCards>
			</WrapContainer>
		);
	}
}

export default MainPage;
