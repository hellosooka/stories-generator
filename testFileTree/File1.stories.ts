import File1 from './File1.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "story/File1",
  component: File1,
  argTypes: { msg: 'string', bebra: 'number[]', bobra: 'number', b: 'boolean', heh: '{ g' },
  parameters: {
    docs: {
      source: {
        code: [
          "import { File1 } from '@estelink/story'\n\n",
          "<File1 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File1 },
    setup: () => ({ args }),
    template: `<File1 v-bind="args" />`,
  });

export const Default = Template.bind({});
