import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const components = [
  ['Button', (await import('@quixsi/components/button')).Button],
  ['Input', (await import('@quixsi/components/input')).Input],
  ['Checkbox', (await import('@quixsi/components/checkbox')).Checkbox],
] as const

const app = createApp(App)

for(const [name, Component] of components) {
  app.component(name, Component)
}

app.use(router)

app.mount('#app')
