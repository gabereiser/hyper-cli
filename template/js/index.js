import "../scss/index.scss";
import * as mdb from 'mdb-ui-kit';
export class Application {
    render() {
        return `
        <div class="container-fluid vh100">
        <span class="logo">⚡️</span>
        </div>
`;
    }
}


const app = new Application();
document.querySelector("#app").innerHTML = app.render();