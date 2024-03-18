<script setup>
import { ref } from 'vue';

import AppMenuItem from './AppMenuItem.vue';

const model = ref([
    {
        label: 'Home',
        items: [{ label: 'Home', to: '/' }]
    },
    {
        label: 'Requests',
        items: [{
            label: 'Requests', icon: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="">
  <path stroke-linecap="round" stroke-linejoin="round" d="M14.25 9.75 16.5 12l-2.25 2.25m-4.5 0L7.5 12l2.25-2.25M6 20.25h12A2.25 2.25 0 0 0 20.25 18V6A2.25 2.25 0 0 0 18 3.75H6A2.25 2.25 0 0 0 3.75 6v12A2.25 2.25 0 0 0 6 20.25Z" />
</svg>
`, to: '/requests',
            size: '20'
        }]
    },
    {
        label: 'UI Components',
        items: [
            { label: 'Form Layout', to: '/uikit/formlayout' },
            { label: 'Input', to: '/uikit/input' },
            { label: 'Float Label', to: '/uikit/floatlabel' },
            { label: 'Invalid State', to: '/uikit/invalidstate' },
            { label: 'Button', class: 'rotated-icon' },
            { label: 'Table', to: '/uikit/table' },
            { label: 'List', to: '/uikit/list' },
            { label: 'Tree', to: '/uikit/tree' },
            { label: 'Panel', to: '/uikit/panel' },
            { label: 'Overlay', to: '/uikit/overlay' },
            { label: 'Media', to: '/uikit/media' },
            { label: 'Menu', preventExact: true },
            { label: 'Message', to: '/uikit/message' },
            { label: 'File', to: '/uikit/file' },
            { label: 'Chart', to: '/uikit/charts' },
            { label: 'Misc', to: '/uikit/misc' }
        ]
    },
    {
        label: 'Prime Blocks',
        items: [
            { label: 'Free Blocks', badge: 'NEW' },
            { label: 'All Blocks', target: '_blank' }
        ]
    },
    {
        label: 'Utilities',
        items: [
            { label: 'PrimeIcons', to: '/utilities/icons' },
            { label: 'PrimeFlex', target: '_blank' }
        ]
    },
    {
        label: 'Pages',

        to: '/pages',
        items: [
            {
                label: 'Landing',

                to: '/landing'
            },
            {
                label: 'Auth',

                items: [
                    {
                        label: 'Login',

                        to: '/auth/login'
                    },
                    {
                        label: 'Error',

                        to: '/auth/error'
                    },
                    {
                        label: 'Access Denied',

                        to: '/auth/access'
                    }
                ]
            },
            {
                label: 'Crud',

                to: '/pages/crud'
            },
            {
                label: 'Timeline',

                to: '/pages/timeline'
            },
            {
                label: 'Not Found',

                to: '/pages/notfound'
            },
            {
                label: 'Empty',

                to: '/pages/empty'
            }
        ]
    },
    {
        label: 'Hierarchy',
        items: [
            {
                label: 'Submenu 1',

                items: [
                    {
                        label: 'Submenu 1.1',

                        items: [
                            { label: 'Submenu 1.1.1', },
                            { label: 'Submenu 1.1.2', },
                            { label: 'Submenu 1.1.3', }
                        ]
                    },
                    {
                        label: 'Submenu 1.2',

                        items: [{ label: 'Submenu 1.2.1', }]
                    }
                ]
            },
            {
                label: 'Submenu 2',

                items: [
                    {
                        label: 'Submenu 2.1',

                        items: [
                            { label: 'Submenu 2.1.1', },
                            { label: 'Submenu 2.1.2', }
                        ]
                    },
                    {
                        label: 'Submenu 2.2',

                        items: [{ label: 'Submenu 2.2.1', }]
                    }
                ]
            }
        ]
    },
    {
        label: 'Get Started',
        items: [
            {
                label: 'Documentation',

                to: '/documentation'
            },
            {
                label: 'Figma',
                url: 'https://www.dropbox.com/scl/fi/bhfwymnk8wu0g5530ceas/sakai-2023.fig?rlkey=u0c8n6xgn44db9t4zkd1brr3l&dl=0',

                target: '_blank'
            },
            {
                label: 'View Source',

                url: 'https://github.com/primefaces/sakai-vue',
                target: '_blank'
            },
            {
                label: 'Nuxt Version',
                url: 'https://github.com/primefaces/sakai-nuxt',
            }
        ]
    }
]);
</script>

<template>
    <ul class="layout-menu">
        <template v-for="(item, i) in model" :key="item">
            <app-menu-item v-if="!item.separator" :item="item" :index="i"></app-menu-item>
            <li v-if="item.separator" class="menu-separator"></li>
        </template>
        <li>
            <a href="https://www.primefaces.org/primeblocks-vue/#/" target="_blank">
                <img src="/layout/images/banner-primeblocks.png" alt="Prime Blocks" class="w-full mt-3" />
            </a>
        </li>
    </ul>
</template>

<style lang="scss" scoped></style>
