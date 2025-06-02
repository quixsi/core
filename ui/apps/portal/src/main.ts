import './assets/main.css'

import { createApp } from 'vue'
import {
  Button,
  Input,
  Checkbox,
  Card,
} from '@quixsi/components'
import App from './App.vue'
import router from './router'

const components = [
  ['Button', Button],
  ['Input', Input],
  ['Checkbox', Checkbox],
  ['Card', Card],
] as const

const app = createApp(App)

for(const [name, Component] of components) {
  app.component(name, Component)
}

app.use(router)

app.mount('#app')
