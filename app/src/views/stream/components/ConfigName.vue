<script setup lang="ts">
import stream from '@/api/stream'
import { message } from 'ant-design-vue'

const props = defineProps<{
  name: string
}>()

const router = useRouter()

const modify = ref(false)
const buffer = ref('')
const loading = ref(false)

watchEffect(() => {
  buffer.value = props.name
})

function clickModify() {
  modify.value = true
}

function save() {
  loading.value = true
  stream.rename(props.name, buffer.value).then(() => {
    modify.value = false
    message.success($gettext('Renamed successfully'))
    router.push({
      path: `/streams/${encodeURIComponent(buffer.value)}`,
    })
  }).finally(() => {
    loading.value = false
  })
}
</script>

<template>
  <div v-if="!modify" class="flex items-center">
    <div class="mr-2">
      {{ buffer }}
    </div>
    <div>
      <AButton type="link" size="small" @click="clickModify">
        {{ $gettext('Rename') }}
      </AButton>
    </div>
  </div>
  <div v-else>
    <AInput v-model:value="buffer">
      <template #suffix>
        <AButton :disabled="buffer === name" type="link" size="small" :loading @click="save">
          {{ $gettext('Save') }}
        </AButton>
      </template>
    </AInput>
  </div>
</template>

<style scoped lang="less">

</style>
