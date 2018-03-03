import React, { Component } from "react";
import styled, { css } from "styled-components";
import Duration from "duration-js";
import moment from "moment";
const Container = styled.div`
	background: #ffffff;
	box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.35);
	padding: 13px;
	display: flex;
	justify-content: space-between;
`;

const WrapField = styled.div`
	font-family: "Quicksand", sans-serif;
	display: grid;
	grid-template-rows: 13px 8px;
	margin-left: 8px;
	margin-right: 8px;
	text-align: center;
`;

const NameFieldText = styled.div`
	font-size: 13px;
	color: #8a8a8a;
	grid-row: 1 / 2;
`;

const ValueFieldText = styled.div`
	font-family: Quicksand-Medium;
	font-size: 16px;
	color: #5b5b5b;
	grid-row: 3 / 4;
`;
const WorkerCardField = props => {
	const nameField = props.name;
	const valueField = props.value;

	return (
		<WrapField>
			<NameFieldText>{nameField}</NameFieldText>
			<ValueFieldText>{valueField}</ValueFieldText>
		</WrapField>
	);
};

const ProgressWork = styled.div`
	margin-top: -4px;
	height: 0px;
	border: 2px solid #60ffeb;
	box-shadow: 0 0 3px 0 #71ffea;
	${props =>
		props.percent &&
		css`
			width: ${(props.percent > 1 ? 1 : props.percent * 100).toFixed(2)}%;
		`};
`;

const SuperContainer = styled.div`
	margin-bottom: 10px;
	margin-top: 10px;
`;

class WorkerCard extends Component {
	state = { now: new Date() };

	constructor(props) {
		super(props);
		this.onRefresh = this.props.onRefresh;
		setInterval(() => {
			this.setState({ now: new Date() });
		}, 1000);
	}

	render() {
		const data = this.props.data;
		const duration = new Duration(data.period / 1000000);

		const lastUpdate = new Date(data.last_update);

		const nextUpdate = moment(lastUpdate)
			.add(duration.minutes(), "m")
			.toDate();

		this.state.now.getTime();

		const num = this.state.now.getTime() - lastUpdate.getTime();
		const den = nextUpdate.getTime() - lastUpdate.getTime();

		let percent = num / den;

		if (this.onRefresh && percent > 0.99) {
			percent = 0.0;
			this.onRefresh(percent);
		}

		return (
			<SuperContainer>
				<Container>
					<WorkerCardField name="Name" value={data.name} />
					<WorkerCardField
						name="Period"
						value={duration.toString()}
					/>
					<WorkerCardField name="State" value={data.state} />
					<WorkerCardField
						name="Time to refresh"
						value={`${
							nextUpdate.getHours() < 10 ? "0" : ""
						}${nextUpdate.getHours()}:${
							nextUpdate.getMinutes() < 10 ? "0" : ""
						}${nextUpdate.getMinutes()}`}
					/>
				</Container>
				<ProgressWork percent={percent} />
			</SuperContainer>
		);
	}
}

export default WorkerCard;
