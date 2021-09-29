import { getAlbums } from "./api"

function App() {
  var send = () => {
    console.log("hello");
    getAlbums();
  };

  return (
    <div className="App">
      <button onClick={send}>Hit</button>
    </div>
  );
}

export default App;
