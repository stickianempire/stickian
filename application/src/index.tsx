import React from 'react';
import ReactDOM from 'react-dom/client';

import App from 'stickian-app/App';
import reportWebVitals from 'reportWebVitals';

import './index.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement,
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);

// Development setup with web vitals for performance measuring
// will be yeeted in the future :)
reportWebVitals(console.log);
