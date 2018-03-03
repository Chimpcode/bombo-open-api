import { observable } from "mobx";
import Duration from "duration-js";
import axios from "axios";

class WorkersStore {
	@observable workers = [];
	@observable creatingNewTask = false;

	constructor() {
		axios
			.get("http://127.0.0.1:8500/manager/get-works", {
				auth: {
					username: "bregymr",
					password: "malpartida1"
				}
			})
			.then(response => {
				this.workers = response.data.data;
			});
	}

	createNewWork({ url, name, type, period }) {
		this.creatingNewTask = true;
		const duration = new Duration(period);
		axios({
			method: "post",
			url: "http://127.0.0.1:8500/manager/add-work",
			auth: {
				username: "bregymr",
				password: "malpartida1"
			},
			headers: {
				"Content-Type": "application/json"
			},
			data: {
				url: url,
				name: name,
				period: duration.milliseconds() * 1000000,
				type: type,
				state: "created"
			}
		})
			.then(response => {
				console.log(response.data.data);

				this.workers = response.data.data;
				this.creatingNewTask = false;
			})
			.catch(response => {
				console.log(response);
			});
	}
}

let store = new WorkersStore();

export default store;
