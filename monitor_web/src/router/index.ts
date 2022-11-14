import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '/@/views/Layout/index.vue'

Vue.use(VueRouter)

export const constantRoutes = [
  {
    path: '/',
    name: 'MainPage',
    meta: {
      icon: 'el-icon-setting',
      title: '首页',
    },
    component: Layout,
    redirect: '/home',
    children: [
      {
        path: '/home',
        name: 'HomePage',
        meta: {
          title: '首页',
        },
        component: () =>
          import(
            '@/views/HomePage/index.vue'
            ),
      },
      {
        path: '/audit',
        name: 'AuditPage',
        meta: {
          title: '审核',
        },
        component: () =>
          import(
            '@/views/HomePage/audit.vue'
            ),
      },
    ],
  },
]

const createRouter = () =>
  new VueRouter({
    routes: constantRoutes,
  })

const router = createRouter()

export function useRouter() {
  return router
}

export function useRoute() {
  return router.currentRoute
}

export default router
