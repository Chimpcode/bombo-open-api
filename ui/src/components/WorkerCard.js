import React, { Component } from "react";
import styled from "styled-components";

const Container = styled.div`
	background: #ffffff;
	box-shadow: 0 2px 6px 0 rgba(0, 0, 0, 0.35);
	padding: 13px;
`;

const WrapField = styled.div`
	font-family: "Quicksand", sans-serif;
`;

const WorkerCardField = props => {
	return;
};

class WorkerCard extends Component {
	constructor(props) {
		super(props);
	}

	render() {
		return <Container />;
	}
}

export default WorkerCard;
