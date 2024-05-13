import { createRouter, createWebHistory } from 'vue-router'
import Home from './components/Home.vue'
import Settings from './components/Settings.vue'
import RssSettings from './components/RssSettings.vue'
import PreferenceSettings from './components/PreferenceSettings.vue'
import AboutSettings from './components/AboutSettings.vue'

const routes = [
  { 
    path: '/', 
    component: Home 
  },
  { 
    path: '/settings', 
    component: Settings,
    redirect: '/settings/rss',
    children: [
      {
        path: 'rss',
        component: RssSettings
      },
      {
        path: 'preference',
        component: PreferenceSettings
      },
      {
        path: 'about',
        component: AboutSettings
      }
    ]
  }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router