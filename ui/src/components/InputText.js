import React, { Component } from "react";
import styled from "styled-components";

const WrapInput = styled.input`
	font-family: "Quicksand", sans-serif;
	color: #454545;
	font-size: 14px;
	padding: 10px;
	border: 0 transparent;
	background: #ffffff;
	box-shadow: 1px 1px 7px 0 rgba(0, 0, 0, 0.18);
	:focus {
		border: 0 transparent;
		outline-offset: 0px !important;
		outline: none !important;
	}
`;

const WrapContainer = styled.div`
	font-family: "Quicksand", sans-serif;
	margin: 10px;
`;

const NameText = styled.span`
	color: #5b5b5b;
	display: block;
	margin-bottom: 13px;
`;

class InputText extends Component {
	constructor(props) {
		super(props);
		this.name = props.name;
		this.onChangeExt = this.props.onChange;
		// this.externalStyle = props.style;
		this.onChange = this.onChange.bind(this);
	}

	onChange(event) {
		if (this.onChangeExt) {
			this.onChangeExt(event);
		}
	}

	render() {
		return (
			<WrapContainer>
				<NameText>{this.name}</NameText>
				<WrapInput style={{ width: "100%" }} onChange={this.onChange} />
			</WrapContainer>
		);
	}
}

export default InputText;
