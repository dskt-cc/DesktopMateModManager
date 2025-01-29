import { TbMelon } from "react-icons/tb";
import { BiPackage, BiSolidPackage } from "react-icons/bi";

export const modTypeButtons = [
    {
        type: "All",
        icon: null,
        label: "All",
        colors:
            "from-miku-aquamarine/90 from-10% via-miku-waterleaf/80 via-50% to-miku-teal/90 to-90%",
    },
    {
        type: "MelonLoader",
        icon: <TbMelon className="text-lg" />,
        label: "MelonLoader",
        colors:
            "from-miku-aquamarine/90 from-10% via-miku-waterleaf/80 via-50% to-miku-teal/90 to-90%",
    },
    {
        type: "BepInEx",
        icon: <BiPackage className="text-lg" />,
        label: "BepInEx",
        colors:
            "from-miku-aquamarine/90 from-10% via-miku-waterleaf/80 via-50% to-miku-teal/90 to-90%",
    },
    {
        type: "Both",
        icon: <BiSolidPackage className="text-lg" />,
        label: "Both",
        colors:
            "from-miku-aquamarine/90 from-10% via-miku-waterleaf/80 via-50% to-miku-teal/90 to-90%",
    },
];
