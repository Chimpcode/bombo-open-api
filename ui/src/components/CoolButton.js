import React, { Component } from "react";
import styled from "styled-components";

const CoolWrap = styled.div`
	background-image: linear-gradient(-90deg, #404ee3 0%, #6c21c2 100%);
	box-shadow: 1px 2px 4px 0 rgba(0, 0, 0, 0.26);
	border-radius: 4px;
	min-height: 20px;
	padding-top: 20px;
	padding-bottom: 20px;
	padding-left: 40px;
	padding-right: 40px;

	-webkit-user-select: none; /* Safari */
	-moz-user-select: none; /* Firefox */
	-ms-user-select: none; /* IE10+/Edge */
	user-select: none; /* Standard */

	:hover {
		box-shadow: 1px 2px 8px 0 rgba(0, 0, 0, 0.4);
		transform: translateY(-3px) scale(1.02);
		cursor: pointer;
	}

	:active {
		box-shadow: 1px 2px 8px 0 rgba(0, 0, 0, 0.4);
		transform: translateY(3px) scale(1.02);
		cursor: pointer;
	}
`;

const InnerButton = styled.div`
	font-size: 18px;
	font-family: "Quicksand", sans-serif;
	color: white;

	text-align: center;
	vertical-align: middle;
`;

class CoolButton extends Component {
	constructor(props) {
		super(props);
		this.onClick = this.props.onClick;
	}

	render() {
		return (
			<CoolWrap onClick={this.onClick ? this.onClick : null}>
				<InnerButton>{this.props.children}</InnerButton>
			</CoolWrap>
		);
	}
}

export default CoolButton;
