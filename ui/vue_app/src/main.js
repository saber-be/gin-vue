import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { vue3Debounce } from 'vue-debounce'
createApp(App).use(ElementPlus).directive('debounce', vue3Debounce({ lock: true })).mount('#app')
