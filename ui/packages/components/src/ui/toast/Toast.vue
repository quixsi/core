<script setup lang="ts">
import type { HTMLAttributes } from 'vue'
import { cn } from '#lib/utils'

defineProps<{
    toasts: Array<{
            message: string
            type?: 'success' | 'error'
            key: string | number
            class?: HTMLAttributes['class']
            duration?: number
    }>
}>()

const emit = defineEmits<{
    (e: 'remove', key: string | number): void
}>()

const handleRemove = (key: string | number) => {
    emit('remove', key)
}
</script>

<template>
    <div class="fixed z-50 top-8 right-8 flex flex-col gap-2 items-end">
        <transition-group name="fade" tag="div">
            <div
                v-for="item in toasts"
                :key="item.key"
                data-slot="toast"
                @click="handleRemove(item.key)"
                :class="cn(
                    'px-4 py-2 mb-3 rounded shadow-lg cursor-pointer',
                    item.type === 'error' ? 'bg-red-500 text-white' : 'bg-green-500 text-white',
                    item.class,
                )"
            >
                {{ item.message }}
            </div>
        </transition-group>
    </div>
</template>

<style scoped>
.fade-enter-active, .fade-leave-active {
    transition: opacity 0.3s;
}

.fade-enter-from, .fade-leave-to {
    opacity: 0;
}
</style>
