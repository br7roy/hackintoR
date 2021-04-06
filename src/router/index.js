import Vue from 'vue'
import Router from 'vue-router'
/* Layout */
import Layout from '@/layout'

Vue.use(Router)

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
import Sch from '@/views/sch'
cmps : {
  Layout,
    Sch
}
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: 'Dashboard', icon: 'dashboard' }
    }]
  },

  {
    path: '/cluster',
    component: Layout,
    redirect: '/cluster/balance',
    name: 'Cluster',
    meta: { title: '集群', icon: 'el-icon-s-help' },
    children: [
      {
        path: 'balance',
        name: 'Balance',
        component: () => import('@/views/cluster'),
        meta: { title: '资源分配', icon: 'tree' }
      },
      {
        path: 'flink',
        name: 'Flink',
        component: () => import('@/views/flink'),
        meta: { title: 'Flink工具', icon: 'table' },
        children: [
          {
            path: 'submit',
            component: () => import('@/views/flink/submit'),
            name: 'Submit',
            meta: { title: '提交Job' }
          },
          {
            path: 'stat',
            component: () => import('@/views/flink/stat'),
            name: 'Stat',
            meta: { title: '查询App状态' }
          }
        ]
      }
    ]
  },
  {
    path: '/sch',
    component: Layout,
    redirect: '/sch/bof',
    name: 'Sch',
    meta: { title: '调度增强', icon: 'eye' },
    children: [
      {
        path: 'bof',
        name: 'Bof',
        component: () => import('@/views/sch/bof'),
        meta: { title: '批量task成败/启停' }
      },
      {
        path: 'bss',
        name: 'Bss',
        component: () => import('@/views/sch/bss'),
        meta: { title: '批量job成败/启停' }
      }
    ]
  },

  {
    path: '/bus',
    component: Layout,
    redirect: '/bus/flow1',
    name: 'Bus',
    meta: {
      title: '业务',
      icon: 'nested'
    },
    children: [
      {
        path: 'flow',
        component: () => import('@/views/nested/menu1'), // Parent router-view
        name: 'Flow',
        meta: { title: '客流工具' },
        children: [
          {
            path: 'mf',
            component: () => import('@/views/nested/menu1/menu1-2'),
            name: 'Mf',
            meta: { title: '商场客流' },
            children: [
              {
                path: 'menu1-2-1',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-1'),
                name: 'Menu1-2-1',
                meta: { title: 'batchDaily' }
              },
              {
                path: 'menu1-2-2',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-2'),
                name: 'Menu1-2-2',
                meta: { title: 'batchResult' }
              },
              {
                path: 'menu1-2-3',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-3'),
                name: 'Menu1-2-3',
                meta: { title: 'batchPython' }
              },
              {
                path: 'menu1-2-4',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-4'),
                name: 'Menu1-2-4',
                meta: { title: '月任务' }
              },
              {
                path: 'menu1-2-5',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-5'),
                name: 'Menu1-2-5',
                meta: { title: '数据同步' }
              },
              {
                path: 'menu1-2-6',
                component: () => import('@/views/nested/menu1/menu1-2/menu1-2-6'),
                name: 'Menu1-2-6',
                meta: { title: '运行GPS比率' }
              }
            ]
          }
        ]
      },
      {
        path: 'polytest',
        component: () => import('@/views/nested/menu2/index'), // Parent router-view
        name: 'Polytest',
        meta: { title: '定点围栏' }
      },
      {
        path: 'ssxs',
        component: () => import('@/views/nested/menu3/index'), // Parent router-view
        name: 'Ssxs',
        meta: { title: '上生新所' }
      }
    ]
  },

  {
    path: 'external-link',
    component: Layout,
    name: 'Link',
    meta: {
      title: '链接',
      icon: 'tree'
    },
    children: [
      {
        path: 'http://bd15-21-33-61:10880/cluster',
        meta: { title: '26YARN', icon: 'link' }
      },
      {
        path: 'http://10-90-49-222-jhdxyjd.mob.local:8088/cluster/apps',
        meta: { title: 'yarn-3.3.3', icon: 'link' }
      },
      {
        path: 'http://10.21.131.11:9800',
        meta: { title: 'ElasticHD', icon: 'link' }
      },
      {
        path: 'http://10.21.131.11:8081',
        meta: { title: 'Flink-WebUI', icon: 'link' }
      }
    ]
  },
  {
    path: '/bd',
    component: Layout,
    redirect: '/bd/shell',
    name: 'Link',
    meta: { title: 'Backdoor', icon: 'el-icon-ice-tea' },
    children: [
      {
        path: 'shell',
        name: 'Shell',
        component: () => import('@/views/backdoor/index'),
        meta: { title: 'Shell', icon: 'el-icon-ice-tea' }
      }
    ]
  },
  {
    path: 'http://gitlab.code.mob.com/futx/hackintor/issues/new',
    component: Layout,
    name: 'BugReport',
    meta: {
      title: 'BugReport',
      icon: 'form'
    }
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
