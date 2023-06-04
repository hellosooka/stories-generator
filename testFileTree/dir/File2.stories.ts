import File2 from './File2.vue'
import { Meta, Story } from "@storybook/vue3";
import { defineComponent } from "vue";

export default {
  title: "story/File2",
  component: File2,
  argTypes: {
    bebra:{},
    febra:{},
    goga:{},
    bob:{}},
  parameters: {
    docs: {
      source: {
        code: [
          "import { File2 } from '@estelink/story'\n\n",
          "<File2 />",
        ].join("\n"),
      },
    },
  },
};

const Template: Story = (args) =>
  defineComponent({
    components: { File2 },
    setup: () => ({ args }),
    template: `<File2 v-bind="args" />`,
  });

export const Default = Template.bind({});
