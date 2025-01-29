import { useState, useEffect } from "react";
import { Navigation } from "./components/Navigation";
import { Mods } from "./containers/Mods";
import { Setup } from "./containers/Setup";
import { GetGamePath } from '../wailsjs/go/main/App';

function App() {
    const [gamePath, setGamePath] = useState('');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const checkGamePath = async () => {
            try {
                const path = await GetGamePath();
                setGamePath(path);
            } catch (error) {
                console.error("Error checking game path:", error);
            } finally {
                setLoading(false);
            }
        };

        checkGamePath();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div className="app-container">
            <Navigation />
            <main>
                {gamePath ? <Mods /> : <Setup />}
            </main>
        </div>
    );
}

export default App;
