import React from 'react';
import { SetGamePath } from '../../wailsjs/go/main/App';

export function Navigation({ onRefresh }) {
    const handleSetPath = async () => {
        try {
            await SetGamePath(); // @TODO @FIX : this just crashes atm
            alert('Game path set successfully!');
        } catch (error) {
            console.error('Failed to set game path:', error);
            alert('Failed to set game path. Please try again.');
        }
    };

    return (
        <nav className="bg-miku-gray shadow-lg">
            <div className="max-w-7xl mx-auto px-4">
                <div className="flex justify-between items-center h-16">
                    <div className="flex items-center">
            <span className="text-miku-teal text-2xl font-bold">
              Desktop Mate Mod Manager
            </span>
                    </div>
                    <div className="flex items-center space-x-4">
                        <button
                            onClick={handleSetPath}
                            className="px-4 py-2 rounded-md text-sm font-medium text-miku-light bg-miku-deep hover:bg-miku-teal transition-colors duration-200"
                        >
                            Set Game Path
                        </button>
                        <button
                            onClick={onRefresh}
                            className="px-4 py-2 rounded-md text-sm font-medium text-miku-light bg-miku-deep hover:bg-miku-teal transition-colors duration-200"
                        >
                            Refresh Mods
                        </button>
                    </div>
                </div>
            </div>
        </nav>
    );
}
