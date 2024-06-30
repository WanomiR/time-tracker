import {Link, Outlet} from "react-router-dom";


const App = () => {
	return (
		<div className="container">
			<div className="row">
				<div className="col">
					<h1 className="mt-3">TimeTracker</h1>
				</div>
				<hr className="mb-3"/>
			</div>

			<div className="row">
				<div className="col-md-2">
					<nav>
						<div className="list-group">
							<Link to={"/"} className={"list-group-item list-group-item-action"}>Timer</Link>
							<Link to={"/tasks"} className={"list-group-item list-group-item-action"}>My Tasks</Link>
							<Link to={"/users"} className={"list-group-item list-group-item-action"}>Users</Link>
						</div>
					</nav>
				</div>
				<div className="col-md-10">
					<Outlet />
				</div>
			</div>
		</div>
	)
}

export default App