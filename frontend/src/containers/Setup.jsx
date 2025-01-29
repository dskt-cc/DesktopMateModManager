
export function Setup() {
    const handleSetPath = async () => {
        // @TODO : implementation once done in Nav
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-[80vh]">
            <div className="max-w-md w-full bg-gray-800 rounded-lg p-8 text-center">
                <h1 className="text-2xl font-bold text-miku-teal mb-4">
                    Welcome to Desktop Mate Mod Manager
                </h1>
                <div className="space-y-4">
                    <p>Desktop Mate installation was not found.</p>
                    <p className="text-miku-waterleaf">
                        Please make sure Desktop Mate is installed and set the correct path.
                    </p>
                    <button
                        onClick={handleSetPath}
                        className="bg-miku-deep hover:bg-miku-teal text-white px-6 py-2 rounded-md transition-colors duration-200"
                    >
                        Set Game Path
                    </button>
                </div>
            </div>
        </div>
    );
}
