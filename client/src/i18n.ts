import { initReactI18next } from "react-i18next";
import i18n from "i18next";

import enUS from "./i18n/en-us.json";
import frCA from "./i18n/fr-ca.json";
import en from "./i18n/en.json";
import fr from "./i18n/fr.json";


i18n.use(initReactI18next).init({
    resources: {
        en: {
            translation: en,
        },
        "en-US": {
            translation: enUS,
        },
        fr: {
            translation: fr,
        },
        "fr-CA": {
            translation: frCA,
        },
    },
    lng: "en",
});