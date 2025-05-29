import { createRoot } from 'react-dom/client';
import { StrictMode } from 'react';

import './assets/styles/global.css';
import './index.css';
import App from './App.tsx';
import './i18n.ts';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
